import { SignUpComponent } from "./sign-up.component";

describe('SignUpComponent', () => {

    it('mounts', () => {
        cy.mount(SignUpComponent);
    });

    it('inputs username', () => {
        const component = cy.mount(SignUpComponent);
        component.get('input[formControlName="username"]')
            .type('example')
            .should('have.value', 'example');
    });

    it('inputs email', () => {
        const component = cy.mount(SignUpComponent);
        component.get('input[formControlName="email"]')
            .type('example@gmail.com')
            .should('have.value', 'example@gmail.com');
    });

    it('inputs password', () => {
        const component = cy.mount(SignUpComponent);
        component.get('input[formControlName="password"]')
            .type('123456')
            .should('have.value', '123456');
    });

    it('inputs date', () => {
        const component = cy.mount(SignUpComponent);
        component.get('input[formControlName="dob"]')
            .type('1990-01-01')
            .should('have.value', '1990-01-01');
    });
});