# Masked Language Models
## Word Sense Disambiguation
Takes as input a context, and a fixed inventory of word senses and output the predicted sense. [Polysemous whole word tokens]
### Uses
Role in model interpretability. Shows interesting distributional properties leading to the _one sense per discourse_ 'rule.' \
### Algorithm
Current best algorithm is the 1-nearest-neighbor algorithm using contextual word embeddings. At training, pass sentence in some 
sense-labeled dataset through any contextual embedding (e.g BERT) resulting in contextual embedding for each token. Then for each sense 
_s_ of any word in the corpus, for each _n_ tokens, we average their _n_ contextual representations to produce a contextual ___sense embedding___
We then calculate nearest neighbor token with the cosine.
### Complications
Anisotrophy exists -- the expected cosine similarity of any pair of words in a corpus is high. Cause: cosine measures are dominated by 
a small number of dimensions of the contextual embedding whose values are very different that the others: ___rogue dimensions___. By 
z-scoring the vectors. 
## Output
Outputs contextual embeddings
## Application
Use natural language to prompt the model. Can use for downstream application 'fine-tuning.' Add a special head on top of pretrained models, taking
their output as input. Using labeled data to train additional parameters at the end. Typically, this process freezes the weights
of the pretrained model.
### Sequence Classification
Classify sequence of text with a label -- commonly called text classication. Prepended special [CLS] token to the front of each input sentence.
The output vector in the final layer of the model for the [CLS] token is the input to the classifier head which makes the relevant
decision. Usual softmax, and cross-entropy rigamarole.
### Sequence-Pair Classification
Tasks include Paraphrase detection, logical entailment and discourse coherence. Exactly the same as NSP objective. During finetuning,
pairs of labeled sentences are presented to the model to produce the h outputs for each input token. [CLS] represents the model's view of 
the pair. To perform classication, [CLS] vector is multiplied by a set of learning classication weights and passed through a softmax. 
- Example: NLI [ Natural Language Inference ] a model is presented with a pair of sentences and must classify the relationship.
### Sequence Labeling
Usually for named entity recognition. PER (person), LOC (location), ORG (organization) or GPE (geo-politcal entity).
