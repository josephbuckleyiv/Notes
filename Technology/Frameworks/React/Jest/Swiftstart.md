# Swiftstart
Install jest as a devDependency using 
```
  npm install --save-dev jest
```
and change 
```json
"scripts": {
  "test": "jest"
}
```
## Create sum.test.js
This corresponds to sum.js

## 
Use it(@param1, @param2)  : where param1 is a string description, and param2 is an anonymous function
```js
  it("should add 1 + 2 to equal 3" () => {
  const result = sum(1,2);
  expect(result).toBe(3)
})
```

## Some Keywords
### toBe
### toEqual
### Describe
This is used to bundle various tests together.
### toBeClose
### toBeGreaterThan
### toBeLessThanOrEqual
