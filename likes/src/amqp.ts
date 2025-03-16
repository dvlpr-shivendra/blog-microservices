import * as amqp from "amqplib";
import { Context, propagation } from "@opentelemetry/api";

const MAX_RETRY_COUNT = 3;
const DLQ = "dlq_main";
export const DELIVERY_MODE_PERSISTENT = 2;
export const POST_LIKED_EVENT = "post.liked";

export async function connect(
  user: string,
  password: string,
  host: string,
  port: string
): Promise<{ channel: amqp.Channel; close: () => Promise<void> }> {
  const address = `amqp://${user}:${password}@${host}:${port}`;
  const conn = await amqp.connect(address);
  const channel = await conn.createChannel();

  await channel.assertExchange(POST_LIKED_EVENT, "fanout", { durable: true });

  return { channel, close: () => conn.close() };
}

export async function handleRetry(
  ch: amqp.Channel,
  msg: amqp.ConsumeMessage | null
): Promise<void> {
  if (!msg) return;

  const headers = msg.properties.headers || {};
  let retryCount = headers["x-retry-count"] || 0;
  retryCount++;
  headers["x-retry-count"] = retryCount;

  console.log(
    `Retrying message ${msg.content.toString()}, retry count: ${retryCount}`
  );

  if (retryCount >= MAX_RETRY_COUNT) {
    console.log(`Moving message to DLQ ${DLQ}`);

    ch.sendToQueue(DLQ, msg.content, {
      headers,
      contentType: "application/json",
      deliveryMode: DELIVERY_MODE_PERSISTENT,
    });
    return;
  }

  await new Promise((resolve) => setTimeout(resolve, retryCount * 1000));

  ch.publish(msg.fields.exchange, msg.fields.routingKey, msg.content, {
    headers,
    contentType: "application/json",
    deliveryMode: DELIVERY_MODE_PERSISTENT,
  });
}

export async function createDLQAndDLX(ch: amqp.Channel): Promise<void> {
  const q = await ch.assertQueue("main_queue", { durable: true });

  const dlx = "dlx_main";

  await ch.assertExchange(dlx, "fanout", { durable: true });

  await ch.bindQueue(q.queue, dlx, "");

  await ch.assertQueue(DLQ, { durable: true });
}

type AmqpHeaderCarrier = Record<string, string>;

export function injectAMQPHeaders(ctx: Context): Record<string, string> {
  const carrier: AmqpHeaderCarrier = {};
  propagation.inject(ctx, carrier);
  return carrier;
}
