# Prompting
## Terminology
### Prompts
Instructions given to a language model.
### Demonstrations
Example to elucidate prompt/instruction.
### In-context Learning
Learning that improves model performance or reduces some loss but does not involve gradient-based updates to the model's performance.
### Instruction-tuning
### Preference Alignment
Often implemented with __RLHF__ or __DPO__ algorithms. A separate model is trained to decide how much a candidate response aligns with human preferences.
### Few Shot Prompting
One or more demonstations. Diminishing returns after first demonstrations.
### Zero-Shot Prompting
When you have zero demonstrations
### Induction Heads
Name for a circuit. Which is part of a neural network, attention computation in transformers.
### Pattern Completion Rule
If the induction head sees B after A each time, it makes the A --> AB
## Transformation
1. Develop a task-specific _template_ that has a free parameter for the input text.
2. Give that input and the task-specific template, the input is used to instantiate a filled prompty that is then passed to a pretrained model.
3. Auto-regressive decoding is then used to generate a sequence of token outputs.
4. The output of the model can either be used directly as the desired output.

## Learning
Model is learning while it processes prompt. In-Context learning is a proposed term to describe the reduced loss at forward pass during inference time of a model without any gradient-based updates.
New research proposes a fuzzy version of the Pattern Completion Rule might be the reason for In-Context Learning. EVIDENCE: Ablating induction heads causes In-Context Learning to decrease.
