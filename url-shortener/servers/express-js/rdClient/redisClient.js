import {createClient} from "redis";

const client = createClient({
    password: process.env.REDIS_PASS,
    socket: {
        host: process.env.REDIS_URI,
        port: process.env.REDIS_PORT,
        tls: {rejectUnauthorized: false},
    }
});

client.on('error', (err) => console.error('Redis Client Error', err));

client.connect().catch((err) => console.error('Redis Connection Error', err));

export default client;