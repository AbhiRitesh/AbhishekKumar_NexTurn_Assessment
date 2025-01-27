const { toggleVisibility } = require("../Exercise5/script");

describe("toggleVisibility", () => {
  let element;
  beforeEach(() => {
    element = { style: { display: "block" } };
  });

  it("should hide the element if it is visible", () => {
    element.style.display = "block";
    toggleVisibility(element);
    expect(element.style.display).toBe("none");
  });

  it("should show the element if it is hidden", () => {
    element.style.display = "none";
    toggleVisibility(element);
    expect(element.style.display).toBe("block");
  });

  it("should call the setter for display", () => {
    spyOn(element.style, "display", "set").and.callFake((value) => {
      this._display = value;
    });
    toggleVisibility(element);
    expect(element.style.display).toBeDefined();
  });
});
