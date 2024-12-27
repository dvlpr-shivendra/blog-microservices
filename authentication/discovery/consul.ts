import Consul from "consul";
import { IRegistery } from "./iRegistery";
import { env } from "../helpers/app";

export class ConsulRegistry implements IRegistery {
    private client: Consul;
    private serviceName: string;
    private host: string;
    private port: number;
    private instanceId: string;
    private healthCheckInterval?: NodeJS.Timeout;

    constructor(serviceHost: string, servicePort: number, serviceName: string) {
        this.serviceName = serviceName;
        this.instanceId = this.generateInstanceID();

        // Service details
        this.host = serviceHost;
        this.port = servicePort;

        // Consul connection details
        const consulHost = env('CONSUL_HOST', 'localhost');
        const consulPort = parseInt(env('CONSUL_PORT', '8500'), 10);

        this.client = new Consul({
            host: consulHost,
            port: consulPort
        });
    }

    private generateInstanceID(): string {
        return `${this.serviceName}-${Date.now()}-${Math.floor(Math.random() * 1000000)}`;
    }

    async register(): Promise<void> {

        await this.client.agent.service.register({
            id: this.instanceId,
            address: this.host,
            port: this.port,
            name: this.serviceName,
            check: {
                name: this.serviceName,
                checkid: this.instanceId,
                tlsskipverify: true,
                ttl: "5s",
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