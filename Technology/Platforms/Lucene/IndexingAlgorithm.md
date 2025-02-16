# Indexing Algorithm

## Benefits.
Index is incrementally updateable. New documents go in tiny new indexes. Bulk of indexes reside in larger 'index segments.'
Lazy merge sort ensures optimal, scalable and n*logn solution. [This happens in the background.]
Supports so-called fast-search. Hits streamed through query processor --> means the memory is proportional to query, not collection.

# Types of Searches Supported
boolean, phrase, range, spatial, span.
