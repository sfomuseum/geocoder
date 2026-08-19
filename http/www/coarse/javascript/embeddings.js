const js_u = new URL(location);
js_u.pathname = js_u.pathname + 'javascript/wllama.js';

const wasm_u = new URL(location);
wasm_u.pathname = wasm_u.pathname + 'wasm/wllama.wasm';

const model_u = new URL(location);
model_u.pathname = model_u.pathname + 'models/bert-bge-small/ggml-model-f16.gguf';

const { Wllama } = await import(js_u.toString());

const CONFIG_PATHS = {
    default: wasm_u.toString(),
};

const wllama = new Wllama(CONFIG_PATHS);
let model_loaded = false;

export async function initModel() {
    
    if (model_loaded) {
	return;
    }
	
    await wllama.loadModelFromUrl(model_u.toString(), {
        embeddings: true,
        pooling_type: 'LLAMA_POOLING_TYPE_MEAN',
        n_ctx: 512,
    });
    
    model_loaded = true;
}

export async function getEmbedding(text) {
    
    if (!model_loaded) {
        throw new Error("Model not initialized. Call initModel() first.");
    }

    if (typeof text !== 'string' || text.trim() === '') {
        throw new TypeError(`Invalid input: Expected non-empty string, got "${typeof text}"`);
    }
    
    const input = text.trim();    
    return await wllama.createEmbedding({input: input});
}
