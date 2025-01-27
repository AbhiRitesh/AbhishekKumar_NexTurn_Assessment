const { getElement } = require("../Exercise2/script");

describe("getElement", () => {
  it("should return the correct element for a valid index", () => {
    const arr = ["a", "b", "c", "d", "e"];
    expect(getElement(arr, 0)).toBe("a");
    expect(getElement(arr, 1)).toBe("b");
    expect(getElement(arr, 2)).toBe("c");
    expect(getElement(arr, 3)).toBe("d");
    expect(getElement(arr, 4)).toBe("e");
  });

  it("should throw an error for negative index", () => {
    const arr = ["a", "b", "c", "d", "e"];
    expect(() => getElement(arr, -1)).toThrowError("Index out of bounds");
  });

  it("should throw error for an index out-of-bounds", () => {
    const arr = ["a", "b", "c", "d", "e"];
    expect(() => getElement(arr, 5)).toThrowError("Index out of bounds");
  });
});
