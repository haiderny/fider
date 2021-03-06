export interface AppSettings {
    mode: string;
    buildTime: string;
    version: string;
    authEndpoint: string;
    environment: string;
    googleAnalytics: string;
    compiler: string;
    domain: string;
}

export interface AuthSettings {
    endpoint: string;
    providers: {
        google: boolean,
        facebook: boolean,
        github: boolean
    };
}
