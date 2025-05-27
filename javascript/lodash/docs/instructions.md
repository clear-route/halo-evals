# Instructions

In this exercise, you'll implement a deep cloning function similar to Lodash's `_.cloneDeep()`. Deep cloning creates a copy of an object or array where all nested objects and arrays are also copied, breaking all references to the original objects.

## Task

The `cloneDeep` function is already defined in the lodash.js file, but it relies on a helper function called `baseClone` which has been replaced with a placeholder that throws an error:

```javascript
function baseClone(value, bitmask, customizer, key, object, stack) {
  throw new Error("Please implement me!")
}
```

Your task is to implement the `baseClone` function to make all the tests pass. The `cloneDeep` function calls `baseClone` with specific flags:

```javascript
function cloneDeep(value) {
  return baseClone(value, CLONE_DEEP_FLAG | CLONE_SYMBOLS_FLAG);
}
```

Where:
- `CLONE_DEEP_FLAG = 1` - Indicates that the clone should be deep (recursive)
- `CLONE_SYMBOLS_FLAG = 4` - Indicates that symbol properties should also be cloned

## Requirements

Your implementation of `baseClone` should handle:

1. **Primitive values**: Numbers, strings, booleans, null, undefined should be returned as is.

2. **Objects and Arrays**: Create new instances and copy all properties/elements.

3. **Deep Cloning**: Recursively clone nested objects and arrays.

4. **Special Objects**: Handle special object types like:
   - Date objects
   - RegExp objects
   - Map objects
   - Set objects
   - Typed arrays (Int8Array, Uint8Array, etc.)
   - ArrayBuffer
   - DataView

5. **Symbol Properties**: Clone properties with Symbol keys.

6. **Functions**: Convert functions to empty objects.

7. **Circular References**: Detect and properly handle circular references to prevent infinite recursion.

## Implementation Hints

1. **Handling Primitives**: Use type checking to return primitive values as is.

2. **Object Type Detection**: Use `Object.prototype.toString.call(value)` or similar methods to detect the type of objects.

3. **Circular References**: Use a `Stack` data structure to track objects being cloned to detect circular references.

4. **Deep Cloning**: Recursively call `baseClone` for nested objects and arrays.

5. **Special Objects**: Create new instances of special objects and copy their properties:
   - For Date: `new Date(original.getTime())`
   - For RegExp: `new RegExp(original.source, original.flags)`
   - For Map/Set: Create new instances and add cloned entries
   - For typed arrays: Create new instances with cloned buffers

## Example

```javascript
// Simple object
const obj = { a: 1, b: { c: 2 } };
const clone = _.cloneDeep(obj);

obj.b.c = 3;
console.log(clone.b.c); // Should be 2, not 3

// Circular reference
const circular = { a: 1 };
circular.self = circular;
const clonedCircular = _.cloneDeep(circular);

console.log(clonedCircular.a); // 1
console.log(clonedCircular.self === clonedCircular); // true
console.log(clonedCircular.self !== circular); // true
```

## Testing

Run the tests with:

```bash
pnpm test
```

The tests cover all the requirements mentioned above, including edge cases and special object types.
