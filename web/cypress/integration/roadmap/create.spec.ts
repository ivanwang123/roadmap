describe("Create Roadmap", () => {
  it("able to create roadmap", () => {
    cy.visit("http://localhost:3000/create/map");

    // main
    cy.get("input[name='title']").type("Roadmap 1");
    cy.get("textarea[name='description']").type("Description");

    // tags
    cy.get("input[name='tag']").type("Tag 1");
    cy.get("[data-testid='add-tag']").click();
    cy.contains("Tag 1").should("exist");

    // checkpoint main
    cy.get("[data-testid='add-checkpoint']").click();
    cy.get("input[name='checkpoints.0.title']").type("Checkpoint 1");
    cy.get("textarea[name='checkpoints.0.instructions']").type("Instructions");

    // links
    cy.get("input[name='link']").type("Link 1");
    cy.get("[data-testid='add-link']").click();
    cy.contains("Link 1").should("exist");

    cy.contains("button", "Create").click();
  });
});
