import { SignUpComponent } from "./sign-up.component";

/**
 * SignUpComponent Cypress test
 */
describe('SignUpComponent', () => {

    // Test that the component mounts
    it('mounts', () => {
        cy.mount(SignUpComponent);
    });

    // Test that username input field is working
    it('inputs username', () => {
        const component = cy.mount(SignUpComponent);
        component.get('input[formControlName="username"]')
            .type('example')
            .should('have.value', 'example');
    });

    // Test that email input field is working
    it('inputs email', () => {
        const component = cy.mount(SignUpComponent);
        component.get('input[formControlName="email"]')
            .type('example@gmail.com')
            .should('have.value', 'example@gmail.com');
    });

    // Test that password input field is working
    it('inputs password', () => {
        const component = cy.mount(SignUpComponent);
        component.get('input[formControlName="password"]')
            .type('123456')
            .should('have.value', '123456');
    });

    // Test that date input field is working
    it('inputs date', () => {
        const component = cy.mount(SignUpComponent);
        component.get('input[formControlName="dob"]')
            .type('1990-01-01')
            .should('have.value', '1990-01-01');
    });
});