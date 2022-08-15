// load type definitions that come with Cypress module
// <reference types="cypress" />

declare namespace Cypress {
  interface Chainable {
    /**
     * Custom command to setup the testing environment.
     */
    setup(): Chainable<Element>;

    /**
     * Custom command to setup the testing environment with fixture data for software and vulnerabilities.
     */
    setupWithSoftware(): Chainable<Element>;

    /**
     * Custom command to login the user programmatically using the fleet API.
     */
    login(email?: string, password?: string): Chainable<Element>;

    /**
     * Custom command to login the user programmatically using the fleet API
     * but with a Cypress session wrapper.
     */
    loginWithCySession(email?: string, password?: string): Chainable<Element>;

    /**
     * Custom command to log out the current user.
     */
    logout(): Chainable<Element>;

    /**
     * Custom command to set a Fleet Desktop token to a host.
     */
    setDesktopToken(hostId?: number, token?: string): Chainable<Element>;

    /**
     * Custom command to add new queries by default.
     */
    seedQueries(): Chainable<Element>;

    /**
     * Custom command to add new scheduled queries by default.
     */
    seedSchedule(): Chainable<Element>;

    /**
     * Custom command to add new policies by default.
     */
    seedPolicies(teamName?: string): Chainable<Element>;

    /**
     * Custom command to add a new user in Fleet (via fleetctl).
     */
    addUser(options?: {
      email?: string;
      password?: string;
      globalRole?: string;
    }): Chainable<Element>;

    /**
     * Custom command to setup the SMTP configuration for this testing environment.
     *
     * NOTE: login() command is required before this, as it will make authenticated
     * requests.
     */
    setupSMTP(): Chainable<Element>;

    /**
     * Custom command to set up SSO auth with the local server.
     *
     * NOTE: login() command is required before this, as it will make authenticated
     * requests.
     */
    setupSSO(options?: {
      enable_sso_idp_login?: boolean;
      enable_jit_provisioning?: boolean;
    }): Chainable<Element>;

    /**
     * Custom command to login a user1@example.com via SSO.
     */
    loginSSO(options?: {
      username?: string;
      password?: string;
    }): Chainable<Element>;

    /**
     * Custom command to get the emails handled by the Mailhog server.
     */
    getEmails(): Chainable;

    /**
     * Custom command to seed the Free tier teams/users.
     *
     * NOTE: login() command is required before this, as it will make authenticated
     * requests.
     */
    seedFree(): Chainable<Element>;

    /**
     * Custom command to seed the Premium tier teams/users.
     *
     * NOTE: login() command is required before this, as it will make authenticated
     * requests.
     */
    seedPremium(): Chainable<Element>;

    /**
     * Custom command to seed the teams/users as represented in Figma.
     *
     * NOTE: login() command is required before this, as it will make authenticated
     * requests.
     */
    seedFigma(): Chainable<Element>;

    /**
     * Custom command to add Docker osquery host.
     *
     * NOTE: login() command is required before this, as it will make authenticated
     * requests.
     */
    addDockerHost(teamName?: string): Chainable;

    /**
     * Custom command to stop any running Docker hosts.
     */
    stopDockerHost(): Chainable;

    /**
     * Custom command to clear downloaded files from test machine.
     */
    clearDownloads(): Chainable;

    /**
     * Custom command to get any element describe only if it is attached to the DOM.
     */
    getAttached(selector: string): Chainable;
  }
}
