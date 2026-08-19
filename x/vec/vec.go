package vec

const DEFAULT_EMBEDDINGS_DIMENSIONS int = 384

// Because this is what https://github.ngxson.com/wllama/examples/basic/ uses
// and it is small. This decision may be revisited.
const DEFAULT_EMBEDDINGS_MODEL string = "hf.co/unsloth/bge-small-en-v1.5-GGUF:F16"

const DEFAULT_EMBEDDER_URI string = "ollama://"
