interface IRegistery {
    register(): Promise<void>;
    deregister(): Promise<void>;
    healthCheck(): Promise<void>;
}

export { IRegistery };