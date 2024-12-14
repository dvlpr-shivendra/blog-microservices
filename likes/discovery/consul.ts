import Consul from "consul";
import { IRegistery } from "./iRegistery";
import { env } from "../helpers/app";

export class ConsulRegistry implements IRegistery {
    private client: Consul;
    private serviceName: string;
    private address: string;
    private instanceId: string;
    private healthCheckInterval?: NodeJS.Timeout;

    constructor(address: string, serviceName: string) {
        this.serviceName = serviceName;
        this.address = address;
        this.instanceId = this.generateInstanceID();

        const [host, portStr] = address.split(':');
        const port = parseInt(portStr, 10);

        this.client = new Consul({
            host: host,
            port: port
        });
    }

    private generateInstanceID(): string {
        return `${this.serviceName}-${Date.now()}-${Math.floor(Math.random() * 1000000)}`;
    }

    async register(): Promise<void> {
        const host = env('GRPC_HOST', 'locslhost');
        const port = parseInt(env('GRPC_PORT', '50051'), 10);

        await this.client.agent.service.register({
            id: this.instanceId,
            address: host,
            port: port,
            name: this.serviceName,
            check: {
                name: this.serviceName,
                checkid: this.instanceId,
                tlsskipverify: true,
                ttl: "5000s",
                timeout: "1s",
                deregistercriticalserviceafter: "10s",
            }
        });

        // Start health check interval
        this.healthCheckInterval = setInterval(() => {
            this.healthCheck().catch(err => {
                console.error('Health check failed:', err);
            });
        }, 1000);
    }

    async deregister(): Promise<void> {
        if (this.healthCheckInterval) {
            clearInterval(this.healthCheckInterval);
        }

        try {
            await this.client.agent.service.deregister(this.instanceId);
        } catch (error) {
            console.error('Failed to deregister service:', error);
            throw error;
        }
    }

    async healthCheck(): Promise<void> {
        try {
            await this.client.agent.check.pass({ 
                id: this.instanceId, 
                note: "online" 
            });
        } catch (error) {
            console.error('Health check failed:', error);
            throw error;
        }
    }
}