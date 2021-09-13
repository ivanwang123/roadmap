describe("Roadmap", () => {
  beforeEach(() => {
    cy.visit("http://localhost:3000/login");
    cy.get("input[name='email']").type("test@test.com");
    cy.get("input[name='password']").type("password");
    cy.contains("button", "Log in").click();
    cy.getCookie("user").then(() => {
      Cypress.Cookies.preserveOnce("user");
    });
  });

  it.skip("able to toggle follow", () => {
    cy.visit("http://localhost:3000/map/1");
    cy.get("[data-testid='unfollow-btn']").should("exist").click();
    cy.get("[data-testid='follow-btn']").should("exist").click();
    cy.get("[data-testid='unfollow-btn']").should("exist");
  });

  it("able to check checkpoint", () => {
    cy.visit("http://localhost:3000/map/1");
    cy.get("[data-testid='check-btn']").first().click();
    cy.get("[data-testid='check-btn']")
      .first()
      .should("have.class", "text-emerald-600");
    cy.get("[data-testid='check-svg']").should("exist");
    cy.get("[data-testid='check-btn']").first().click();
    cy.get("[data-testid='check-svg']").should("not.exist");
  });
});
