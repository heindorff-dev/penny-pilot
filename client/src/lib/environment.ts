class Environment {
    PREFIX_VALUE: string;
    constructor() {
        this.PREFIX_VALUE = import.meta.env.VITE_PREFIX_VALUE;
    }
}

const environment = new Environment();

export default environment;