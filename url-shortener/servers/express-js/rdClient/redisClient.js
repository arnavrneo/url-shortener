import {createClient} from "redis";

const client = createClient({
    password: process.env.REDIS_PASS,
    socket: {
        host: 'redis-15391.c326.us-east-1-3.ec2.cloud.redislabs.com',
        port: 15391,
        tls: {rejectUnauthorized: false},
    }
});

client.on('error', (err) => console.error('Redis Client Error', err));

client.connect().catch((err) => console.error('Redis Connection Error', err));

export default client;