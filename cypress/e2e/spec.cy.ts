describe('template spec', () => {
  
  it('top games should display games', () => {
    cy.visit('localhost:4200')

    cy.contains('Top Games').click()

    cy.contains('Grand Theft Auto V')
  });

  it('should sign up', () => {
    cy.visit('localhost:4200')

    cy.contains('Sign In').click()

    cy.contains('Sign Up').click()
   
    cy.get('input[formControlName="username"]').type('admin', {force: true})
    cy.get('input[formControlName="email"]').type('admin@localhost', {force: true})
    cy.get('input[formControlName="password"]').type('password', {force: true})
  });

  it('should sign in', () => {
    cy.visit('localhost:4200')

    cy.contains('Sign In').click()

    cy.get('input[formControlName="username"]').type('admin', {force: true})
    cy.get('input[formControlName="password"]').type('password', {force: true})
  });

});