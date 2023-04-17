describe('template spec', () => {

  /**
   * Clicking About Us should bring you to about us page
   */
  it('should go to about us page', () => {
    cy.visit('localhost:4200')

    cy.get('button').contains('About Us').click()

    cy.url().should('include', '/about-us')
  });
  
  /**
   * Test that the top games page displays games
   */
  it('top games should display games', () => {
    cy.visit('localhost:4200')

    cy.contains('Top Games').click()

    cy.contains('Grand Theft Auto V')
  });

  /**
   * Test that the top games page buttons work
   */
  it('top games buttons should work', () => {

    cy.visit('localhost:4200/top-games')

    cy.get('button').contains('Next 40 >').click()

    cy.contains('Apex Legends')

    cy.get('button').contains('< Previous 40').click()

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

    cy.get('button').contains('Sign Up').click()

    cy.on('window:alert', (str) => {
      expect(str).to.equal('This user already exists');
    });

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

    cy.get('img[class="profile-image"]').click()
    cy.get('button').contains('Profile').click()
    cy.url().should('include', '/user')
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