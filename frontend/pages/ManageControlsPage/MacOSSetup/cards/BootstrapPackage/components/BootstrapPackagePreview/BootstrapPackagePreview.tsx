import React from "react";

const baseClass = "bootstrap-package-preview";

const BootstrapPackagePreview = () => {
  return (
    <div className={baseClass}>
      <h2>End user experience</h2>
      <p>
        The bootstrap package is automatically installed after the end user
        authenticates and agrees to the EULA during the Remote Management pane
        in macOS Setup Assistant.
      </p>
      <p>
        The end user is allowed to continue to the next setup pane before the
        installation starts.
      </p>
      <p>The package isn&apos;t installed on hosts that already enrolled.</p>
      <img
        className={`${baseClass}__preview-img`}
        src={"test"}
        alt="OS setup preview"
      />
    </div>
  );
};

export default BootstrapPackagePreview;
