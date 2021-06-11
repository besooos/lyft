import React from "react";
import { FormFields } from "@clutch-sh/experimentation";
import { shallow } from "enzyme";

import { ExperimentDetails } from "../start-experiment";

jest.mock("react-router-dom", () => {
  return {
    ...jest.requireActual("react-router-dom"),
    useNavigate: jest.fn(),
  };
});

describe("Start Experiment workflow", () => {
  it("renders correctly", () => {
    const component = shallow(
      <ExperimentDetails upstreamClusterTypeSelectionEnabled={false} onStart={() => {}} />
    );
    expect(component.find(FormFields).dive().debug()).toMatchSnapshot();
  });

  it("renders correctly with upstream cluster type selection enabled", () => {
    const component = shallow(
      <ExperimentDetails upstreamClusterTypeSelectionEnabled onStart={() => {}} />
    );
    expect(component.find(FormFields).dive().debug()).toMatchSnapshot();
  });
});
