import { AppContext } from "context/app";
import React, { useContext } from "react";
import { Params } from "react-router/lib/Router";

import SideNav from "../components/SideNav";
import INTEGRATION_SETTINGS_NAV_ITEMS from "./IntegrationNavItems";

const baseClass = "integrations";

interface IIntegrationSettingsPageProps {
  params: Params;
}

const IntegrationsPage = ({ params }: IIntegrationSettingsPageProps) => {
  const { section } = params;
  const DEFAULT_SETTINGS_SECTION = INTEGRATION_SETTINGS_NAV_ITEMS[0];
  const navItems = INTEGRATION_SETTINGS_NAV_ITEMS;
  const currentSection =
    navItems.find((item) => item.urlSection === section) ??
    DEFAULT_SETTINGS_SECTION;

  const CurrentCard = currentSection.Card;

  return (
    <div className={`${baseClass}`}>
      <p className={`${baseClass}__page-description`}>
        Add ticket destinations and turn on mobile device management features.
      </p>
      <SideNav
        className={`${baseClass}__side-nav`}
        navItems={navItems}
        activeItem={currentSection.urlSection}
        CurrentCard={<CurrentCard />}
      />
    </div>
  );
};

export default IntegrationsPage;
