import { Elysia } from "elysia";
import { opentelemetry } from '@elysiajs/opentelemetry'
import { swagger } from '@elysiajs/swagger'

import { controller } from "./controllers";

import { BatchSpanProcessor } from '@opentelemetry/sdk-trace-node'
import { OTLPTraceExporter } from '@opentelemetry/exporter-trace-otlp-proto'

export const app = new Elysia()
    .use(opentelemetry({
        spanProcessors: [
            new BatchSpanProcessor(
                new OTLPTraceExporter()
            )
        ]
    })) 
    .use(swagger({
        scalarConfig: {
            defaultHttpClient: {
                targetKey: 'javascript',
                clientKey: 'axios',
            },
        }
    }))
    .onError(({ error, code }) => { 
        if (code === 'NOT_FOUND') return 'Not Found :('

        console.error(error) 
    }) 
    // .use(base)
    // .use(user)
    // .use(note)
    // .use(date)
    .use(controller)
    .listen(3000);

console.log(
  `🦊 Elysia is running at ${app.server?.hostname}:${app.server?.port}`
);
