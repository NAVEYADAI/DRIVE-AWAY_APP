/// <reference types="vite/client" />

interface ImportMetaEnv {
    // Mapbox *public* token (pk.*), restricted by URL in the Mapbox dashboard.
    readonly VITE_MAPBOX_TOKEN?: string;
}
