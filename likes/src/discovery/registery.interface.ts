interface Registery {
    register(): Promise<void>;
    deregister(): Promise<void>;
    healthCheck(): Promise<void>;
}

export { Registery };