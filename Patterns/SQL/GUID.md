## Overhead
Working with GUID inside SQL databases -- joining, querying with, or manipulating GUIDs can significantly impact performance and readability. For that reason, recommend creating a table mapping objects with GUIDs to auto-increment primary key, which can then be used as key for accessing other tables.
