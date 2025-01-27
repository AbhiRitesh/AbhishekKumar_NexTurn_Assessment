const { sendNotification } = require("../Exercise4/script");

describe("sendNotification", () => {
  it("should return 'Notification Sent' when the notification is sent", () => {
    const notificationService = {
      send: (message) => true,
    };
    expect(sendNotification(notificationService, "Hello")).toBe(
      "Notification Sent"
    );
  });

  it("should return 'Failed to Send' when the notification is not sent", () => {
    const notificationService = {
      send: (message) => false,
    };
    expect(sendNotification(notificationService, "Hello")).toBe(
      "Failed to Send"
    );
  });
});
