// 1. Import the locally hosted JS entry point
import { Wllama } from '/javascript/wllama/esm/index.js';

// 2. Point to your locally hosted WASM binary locations
const CONFIG_PATHS = {
    default: '/wasm/wllama.wasm',
};

const MODEL_URL = '/models/bert-bge-small/ggml-model-f16.gguf';

const wllama = new Wllama(CONFIG_PATHS);
let isModelLoaded = false;

export async function getEmbedding(text) {
  if (!isModelLoaded) {
    await wllama.loadModelFromUrl(MODEL_URL, {
      embeddings: true,
    });
    isModelLoaded = true;
  }

  return await wllama.createEmbedding(text);
}
