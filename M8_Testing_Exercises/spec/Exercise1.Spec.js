const { capitalize, reverseString } = require("../Exercise1/Capitalize.js");

describe("capitalize", () => {
  it("should capatalize the first letter of the word", () => {
    expect(capitalize("hello")).toBe("Hello");
  });
  it("should return an empty when the input is an empty string", () => {
    expect(capitalize("")).toBe("");
  });
  it("should return single char when input is single character", () => {
    expect(capitalize("a")).toBe("A");
  });
  it("should return the same string if already capitalized", () => {
    expect(capitalize("Hello")).toBe("Hello");
  });
});

describe("reverseString", () => {
  it("should reverse a string", () => {
    expect(reverseString("hello")).toBe("olleh");
  });

  it("should return an empty string when input is an empty string", () => {
    expect(reverseString("")).toBe("");
  });

  it("should return a single character when input is a single character", () => {
    expect(reverseString("a")).toBe("a");
  });

  it("should return the same string if palindorme", () => {
    expect(reverseString("madam")).toBe("madam");
  });
});
