describe('template spec', () => {
  
  /**
   * Test that the top games page displays games
   */
  it('top games should display games', () => {
    cy.visit('localhost:4200')

    cy.contains('Top Games').click()

    cy.contains('Grand Theft Auto V')
  });

  /**
   * Test that sign up displays the sign up form and that the form can be filled out
   */
  it('should sign up', () => {
    cy.visit('localhost:4200')

    cy.contains('Sign In').click()

    cy.contains('Sign Up').click()
   
    cy.get('input[formControlName="username"]').type('admin', {force: true})
    cy.get('input[formControlName="email"]').type('admin@localhost', {force: true})
    cy.get('input[formControlName="password"]').type('password', {force: true})

    const stub = cy.stub()

    cy.on('window:alert', stub)

    cy.get('button').contains('Sign Up').click()
    .then(() => {
      expect(stub.getCall(0)).to.be.calledWith("This user already exists")
    })

  });

  /**
   * Test that sign in displays the sign in form and that the form can be filled out
   */
  it('should sign in', () => {
    cy.visit('localhost:4200')

    cy.contains('Sign In').click()

    cy.get('input[formControlName="username"]').type('admin', {force: true})
    cy.get('input[formControlName="password"]').type('password', {force: true})

    cy.get('.btn').click()

    cy.url().should('include', '/home')
  });

  /**
   * Clicking a game should bring you to game page
   */
  it('should go to game page on click', () => {
    cy.visit('localhost:4200')

    cy.get('img[class="featured-game"]').click()

    cy.url().should('include', '/game')
  })

});