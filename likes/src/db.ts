import mongoose from "mongoose";
import { env } from "./helpers/app";

// Connection options with pooling configuration
const mongoOptions: mongoose.ConnectOptions = {
  maxPoolSize: parseInt(env("MONGODB_MAX_POOL_SIZE", "10")),
  minPoolSize: parseInt(env("MONGODB_MIN_POOL_SIZE", "5")),
  serverSelectionTimeoutMS: parseInt(
    env("MONGODB_SERVER_SELECTION_TIMEOUT_MS", "5000")
  ),
  socketTimeoutMS: parseInt(env("MONGODB_SOCKET_TIMEOUT_MS", "45000")),
  connectTimeoutMS: parseInt(env("MONGODB_CONNECT_TIMEOUT_MS", "10000")),
  heartbeatFrequencyMS: parseInt(
    env("MONGODB_HEARTBEAT_FREQUENCY_MS", "10000")
  ),
  retryWrites: true,
  retryReads: true,
};

// Singleton pattern for MongoDB connection
class Database {
  private static instance: Database;
  private isConnected = false;

  private constructor() {}

  public static getInstance(): Database {
    if (!Database.instance) {
      Database.instance = new Database();
    }
    return Database.instance;
  }

  public async connect(): Promise<void> {
    if (this.isConnected) {
      console.log("Using existing MongoDB connection");
      return;
    }

    try {
      const uri = env("MONGODB_URI", "");
      if (!uri) {
        throw new Error("MONGODB_URI environment variable is not set");
      }

      // Set up connection monitoring
      mongoose.connection.on("connected", () => {
        console.log("MongoDB connection established");
        this.isConnected = true;
      });

      mongoose.connection.on("error", (err) => {
        console.error("MongoDB connection error:", err);
        this.isConnected = false;
      });

      mongoose.connection.on("disconnected", () => {
        console.log("MongoDB disconnected");
        this.isConnected = false;
      });

      // Handle application termination
      process.on("SIGINT", async () => {
        await mongoose.connection.close();
        console.log("MongoDB connection closed due to app termination");
        process.exit(0);
      });

      // Connect with pooling options
      await mongoose.connect(uri, mongoOptions);
    } catch (error) {
      console.error("Failed to connect to MongoDB:", error);
      throw error;
    }
  }

  public async disconnect(): Promise<void> {
    if (!this.isConnected) return;

    try {
      await mongoose.connection.close();
      this.isConnected = false;
      console.log("MongoDB disconnected successfully");
    } catch (error) {
      console.error("Error while disconnecting from MongoDB:", error);
      throw error;
    }
  }

  public getConnection(): mongoose.Connection {
    return mongoose.connection;
  }
}

// Export the singleton instance
export const db = Database.getInstance();
