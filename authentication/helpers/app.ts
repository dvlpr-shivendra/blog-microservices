export function env(key: string, fallback: string) {
    return process.env[key] || fallback;
}
