import { defineConfig } from "astro/config";
import tailwind from "@astrojs/tailwind";
import vercel from "@astrojs/vercel/serverless";

import react from "@astrojs/react";

// https://astro.build/config
export default defineConfig({
    output: "server",
    adapter: vercel(),
    integrations: [
        tailwind({
            applyBaseStyles: false
        }),
        react()
    ],
    markdown: {
        shikiConfig: {
            theme: "dracula",
            wrap: true
        }
    }
});
