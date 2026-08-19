import { Wllama } from '/javascript/wllama.js';

const CONFIG_PATHS = {
    default: '/wasm/wllama.wasm',
};

const MODEL_URL = '/models/bert-bge-small/ggml-model-f16.gguf';

const wllama = new Wllama(CONFIG_PATHS);
let isModelLoaded = false;

export async function initModel() {
    if (isModelLoaded) return;
    
    await wllama.loadModelFromUrl(MODEL_URL, {
        // 💡 Ensure embeddings is plural 'embeddings: true'
        embeddings: true,
        pooling_type: 'LLAMA_POOLING_TYPE_MEAN',
        n_ctx: 512,
        useCache: false, // Prevents OPFS filesystem permission errors
    });
    
    isModelLoaded = true;
}

export async function getEmbedding(text) {
    if (!isModelLoaded) {
        throw new Error("Model not initialized. Call initModel() first.");
    }

    if (typeof text !== 'string' || text.trim() === '') {
        throw new TypeError(`Invalid input: Expected non-empty string, got "${typeof text}"`);
    }
    
    const safeText = text.trim();
    
    // 💡 Pass the clean string primitive directly. 
    // It returns the flat array of floating-point numbers automatically.
    return await wllama.createEmbedding({input: safeText});
}
