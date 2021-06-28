describe("Login test", () => {
  it("Logins in user", () => {
    cy.visit("http://localhost:3000/login");

    cy.get("#email")
      .type("test@test.com")
      .should("have.value", "test@test.com");

    cy.get("#password").type("password").should("have.value", "password");

    cy.get("button[type='submit']").click();

    cy.location("pathname").should("eq", "/");

    cy.getCookie("user").should("exist");
    cy.getCookie("refresh").should("exist");
  });
});
