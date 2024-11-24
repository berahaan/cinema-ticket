// nuxt.d.ts
import { NuxtConfig } from "nuxt";

declare module "nuxt" {
  interface NuxtConfig {
    apollo?: {
      clients?: {
        [key: string]: {
          httpEndpoint?: string;
          wsEndpoint?: string;
          link?: (options: { httpEndpoint: string; wsEndpoint: string }) => any;
          tokenName?: string;
          defaultOptions?: {
            $query?: {
              fetchPolicy?: string;
            };
          };
        };
      };
    };
  }
}
