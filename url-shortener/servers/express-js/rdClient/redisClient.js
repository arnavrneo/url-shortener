import {createClient} from "redis";

const REDIS_PASS = process.env.REDIS_PASS;
const REDIS_URI = process.env.REDIS_URI;
const REDIS_PORT = process.env.REDIS_PORT;

const client = createClient({
    password: REDIS_PASS,
    socket: {
        host: REDIS_URI,
        port: REDIS_PORT,
        tls: {rejectUnauthorized: false},
    }
});

client.on('error', (err) => console.error('Redis Client Error', err));

client.connect().catch((err) => console.error('Redis Connection Error', err));

export default client;