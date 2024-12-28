"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.ConsulRegistry = void 0;
const consul_1 = __importDefault(require("consul"));
const app_1 = require("../helpers/app");
class ConsulRegistry {
    client;
    serviceName;
    host;
    port;
    instanceId;
    healthCheckInterval;
    constructor(serviceHost, servicePort, serviceName) {
        this.serviceName = serviceName;
        this.instanceId = this.generateInstanceID();
        // Service details
        this.host = serviceHost;
        this.port = servicePort;
        // Consul connection details
        const consulHost = (0, app_1.env)('CONSUL_HOST', 'localhost');
        const consulPort = parseInt((0, app_1.env)('CONSUL_PORT', '8500'), 10);
        this.client = new consul_1.default({
            host: consulHost,
            port: consulPort
        });
    }
    generateInstanceID() {
        return `${this.serviceName}-${Date.now()}-${Math.floor(Math.random() * 1000000)}`;
    }
    async register() {
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
    async deregister() {
        if (this.healthCheckInterval) {
            clearInterval(this.healthCheckInterval);
        }
        try {
            await this.client.agent.service.deregister(this.instanceId);
        }
        catch (error) {
            console.error('Failed to deregister service:', error);
            throw error;
        }
    }
    async healthCheck() {
        try {
            await this.client.agent.check.pass({
                id: this.instanceId,
                note: "online"
            });
        }
        catch (error) {
            console.error('Health check failed:', error);
            throw error;
        }
    }
}
exports.ConsulRegistry = ConsulRegistry;
