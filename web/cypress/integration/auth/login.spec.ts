describe("Login", () => {
  it("able to log in", () => {
    cy.visit("http://localhost:3000/login");
    cy.get("input[name='email']").type("test@test.com");
    cy.get("input[name='password']").type("password");
    cy.contains("button", "Log in").click();

    cy.location("pathname").should("eq", "/");
    cy.get("[data-testid='alert-notification']")
      .contains("Successfully")
      .should("exist");
    cy.getCookie("user").should("exist");
    cy.getCookie("refresh").should("exist");
  });

  it("errors with empty fields", () => {
    cy.visit("http://localhost:3000/login");
    cy.contains("button", "Log in").click();
    cy.contains("valid email").should("exist");
    cy.contains("required").should("exist");
  });

  it("errors with wrong credentials", () => {
    cy.visit("http://localhost:3000/login");
    cy.get("input[name='email']").type("invalid@test.com");
    cy.get("input[name='password']").type("password");
    cy.contains("button", "Log in").click();

    cy.get("[data-testid='text-notification']")
      .contains("no user")
      .should("exist");
  });
});

export {};
