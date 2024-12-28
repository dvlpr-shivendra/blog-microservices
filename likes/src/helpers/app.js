"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.env = env;
function env(key, fallback) {
    return process.env[key] || fallback;
}
