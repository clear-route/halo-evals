const _ = require("./lodash");

describe("cloneDeep", () => {
  // Primitive values
  test("should clone primitive values", () => {
    expect(_.cloneDeep(null)).toBe(null);
    expect(_.cloneDeep(undefined)).toBe(undefined);
    expect(_.cloneDeep(1)).toBe(1);
    expect(_.cloneDeep("string")).toBe("string");
    expect(_.cloneDeep(true)).toBe(true);
    expect(_.cloneDeep(false)).toBe(false);
    expect(_.cloneDeep(NaN)).toBeNaN();
    expect(_.cloneDeep(Infinity)).toBe(Infinity);
    expect(_.cloneDeep(-Infinity)).toBe(-Infinity);
  });

  // Simple objects
  test("should clone an object", () => {
    const obj = { a: 1, b: 2 };
    const result = _.cloneDeep(obj);
    expect(result).toEqual(obj);
    expect(result).not.toBe(obj);
    obj.a = 3;
    expect(result.a).toBe(1);
  });

  // Nested objects
  test("should deep clone nested objects", () => {
    const obj = { a: 1, b: { c: 2, d: { e: 3 } } };
    const result = _.cloneDeep(obj);

    expect(result).toEqual(obj);
    expect(result).not.toBe(obj);
    expect(result.b).not.toBe(obj.b);
    expect(result.b.d).not.toBe(obj.b.d);

    obj.b.c = 4;
    obj.b.d.e = 5;

    expect(result.b.c).toBe(2);
    expect(result.b.d.e).toBe(3);
  });

  // Arrays
  test("should clone arrays", () => {
    const arr = [1, 2, 3];
    const result = _.cloneDeep(arr);

    expect(result).toEqual(arr);
    expect(result).not.toBe(arr);

    arr[0] = 4;
    expect(result[0]).toBe(1);
  });

  // Nested arrays
  test("should deep clone nested arrays", () => {
    const arr = [1, [2, [3, 4]]];
    const result = _.cloneDeep(arr);

    expect(result).toEqual(arr);
    expect(result[1]).not.toBe(arr[1]);
    expect(result[1][1]).not.toBe(arr[1][1]);

    arr[1][0] = 5;
    arr[1][1][0] = 6;

    expect(result[1][0]).toBe(2);
    expect(result[1][1][0]).toBe(3);
  });

  // Mixed objects and arrays
  test("should clone objects with arrays and arrays with objects", () => {
    const obj = { a: [1, 2, { b: 3 }] };
    const arr = [1, { c: 2, d: [3, 4] }];

    const objResult = _.cloneDeep(obj);
    const arrResult = _.cloneDeep(arr);

    expect(objResult).toEqual(obj);
    expect(objResult.a).not.toBe(obj.a);
    expect(objResult.a[2]).not.toBe(obj.a[2]);

    expect(arrResult).toEqual(arr);
    expect(arrResult[1]).not.toBe(arr[1]);
    expect(arrResult[1].d).not.toBe(arr[1].d);

    obj.a[2].b = 5;
    arr[1].d[0] = 5;

    expect(objResult.a[2].b).toBe(3);
    expect(arrResult[1].d[0]).toBe(3);
  });

  // Date objects
  test("should clone Date objects", () => {
    const date = new Date();
    const result = _.cloneDeep(date);

    expect(result).toEqual(date);
    expect(result).not.toBe(date);
    expect(result.getTime()).toBe(date.getTime());
  });

  // RegExp objects
  test("should clone RegExp objects", () => {
    const regex = /test/gi;
    const result = _.cloneDeep(regex);

    expect(result).not.toBe(regex);
    expect(result.source).toBe(regex.source);
    expect(result.flags).toBe(regex.flags);
    expect(result.test("TEST")).toBe(true);
  });

  // Map objects
  test("should clone Map objects", () => {
    const map = new Map([
      ["a", 1],
      ["b", { c: 2 }],
    ]);
    const result = _.cloneDeep(map);

    expect(result).not.toBe(map);
    expect(result.size).toBe(map.size);
    expect(result.get("a")).toBe(1);
    expect(result.get("b")).toEqual({ c: 2 });
    expect(result.get("b")).not.toBe(map.get("b"));

    map.get("b").c = 3;
    expect(result.get("b").c).toBe(2);
  });

  // Set objects
  test("should clone Set objects", () => {
    const obj = { a: 1 };
    const set = new Set([1, obj, "test"]);
    const result = _.cloneDeep(set);

    expect(result).not.toBe(set);
    expect(result.size).toBe(set.size);
    expect(result.has(1)).toBe(true);
    expect(result.has("test")).toBe(true);

    // The object in the set should be cloned
    const resultObj = Array.from(result).find(
      (item) => typeof item === "object"
    );
    expect(resultObj).toEqual(obj);
    expect(resultObj).not.toBe(obj);

    obj.a = 2;
    expect(resultObj.a).toBe(1);
  });

  // Symbol properties
  test("should clone symbol properties", () => {
    const symbol = Symbol("test");
    const obj = { [symbol]: 1 };
    const result = _.cloneDeep(obj);

    expect(result[symbol]).toBe(1);

    obj[symbol] = 2;
    expect(result[symbol]).toBe(1);
  });

  // Functions
  test("should handle functions", () => {
    const func = function () {
      return 42;
    };
    const result = _.cloneDeep(func);

    // In Lodash's implementation, functions are actually converted to empty objects
    expect(result).toEqual({});
  });

  // Objects with methods
  test("should clone objects with methods", () => {
    const obj = {
      value: 42,
      getValue: function () {
        return this.value;
      },
    };

    const result = _.cloneDeep(obj);

    expect(result).not.toBe(obj);
    expect(result.getValue()).toBe(42);

    obj.value = 100;
    expect(result.getValue()).toBe(42);
  });

  // Circular references
  test("should handle circular references", () => {
    const obj = { a: 1 };
    obj.self = obj;

    const result = _.cloneDeep(obj);

    expect(result).not.toBe(obj);
    expect(result.a).toBe(1);
    expect(result.self).toBe(result);
  });

  // Complex circular references
  test("should handle complex circular references", () => {
    const obj1 = { a: 1 };
    const obj2 = { b: 2, ref: obj1 };
    obj1.ref = obj2;

    const result = _.cloneDeep(obj1);

    expect(result).not.toBe(obj1);
    expect(result.a).toBe(1);
    expect(result.ref).not.toBe(obj2);
    expect(result.ref.b).toBe(2);
    expect(result.ref.ref).toBe(result);
  });

  // Array circular references
  test("should handle array circular references", () => {
    const arr = [1, 2, 3];
    arr.push(arr);

    const result = _.cloneDeep(arr);

    expect(result).not.toBe(arr);
    expect(result[0]).toBe(1);
    expect(result[3]).toBe(result);
  });

  // Typed arrays
  test("should clone typed arrays", () => {
    const typedArrays = [
      new Int8Array([1, 2, 3]),
      new Uint8Array([1, 2, 3]),
      new Uint8ClampedArray([1, 2, 3]),
      new Int16Array([1, 2, 3]),
      new Uint16Array([1, 2, 3]),
      new Int32Array([1, 2, 3]),
      new Uint32Array([1, 2, 3]),
      new Float32Array([1, 2, 3]),
      new Float64Array([1, 2, 3]),
    ];

    typedArrays.forEach((typedArray) => {
      const result = _.cloneDeep(typedArray);

      expect(result).not.toBe(typedArray);
      expect(result).toEqual(typedArray);

      typedArray[0] = 100;
      expect(result[0]).toBe(1);
    });
  });

  // ArrayBuffer
  test("should clone ArrayBuffer", () => {
    const buffer = new ArrayBuffer(8);
    const view = new Uint8Array(buffer);
    view[0] = 1;

    const result = _.cloneDeep(buffer);

    expect(result).not.toBe(buffer);
    expect(result.byteLength).toBe(buffer.byteLength);

    const resultView = new Uint8Array(result);
    expect(resultView[0]).toBe(1);

    view[0] = 2;
    expect(resultView[0]).toBe(1);
  });

  // DataView
  test("should clone DataView", () => {
    const buffer = new ArrayBuffer(8);
    const dataView = new DataView(buffer);
    dataView.setInt16(0, 42);

    const result = _.cloneDeep(dataView);

    expect(result).not.toBe(dataView);
    expect(result.byteLength).toBe(dataView.byteLength);
    expect(result.getInt16(0)).toBe(42);

    dataView.setInt16(0, 100);
    expect(result.getInt16(0)).toBe(42);
  });

  // Edge cases
  test("should handle edge cases", () => {
    // Empty objects and arrays
    expect(_.cloneDeep({})).toEqual({});
    expect(_.cloneDeep([])).toEqual([]);

    // Object with no prototype
    const noProto = Object.create(null);
    noProto.a = 1;
    const result = _.cloneDeep(noProto);

    expect(result).not.toBe(noProto);
    expect(result.a).toBe(1);
    expect(Object.getPrototypeOf(result)).not.toBe(null);
  });
});
