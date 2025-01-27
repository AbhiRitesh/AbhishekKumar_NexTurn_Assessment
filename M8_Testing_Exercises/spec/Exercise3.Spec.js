const { delayedGreeting } = require("../Exercise3/script");

describe("delayedGreeting", () => {
  beforeEach(() => {
    jasmine.clock().install();
  });

  afterEach(() => {
    jasmine.clock().uninstall();
  });

  it("should resolve with the correct greeting message after the delay", (done) => {
    delayedGreeting("John", 2000).then((message) => {
      expect(message).toBe("Hello, John!");
      done();
    });
    jasmine.clock().tick(2000);
  });

  it("should respect the delay before resolving", (done) => {
    let resolved = false;
    delayedGreeting("Jane", 2000).then(() => {
      resolved = true;
      expect(resolved).toBe(true);
      done();
    });
    jasmine.clock().tick(1999);
    expect(resolved).toBe(false);
    jasmine.clock().tick(1);
  });
});
