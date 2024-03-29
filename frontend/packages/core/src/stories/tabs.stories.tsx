import * as React from "react";
import type { Meta } from "@storybook/react";

import type { TabsProps } from "../tab";
import { Tab, Tabs } from "../tab";

export default {
  title: "Core/Tab/Tabs",
  component: Tab,
  argTypes: {
    onClick: { action: "onClick event" },
  },
} as Meta;

const Template = ({ tabCount, value, variant }: TabsProps & { tabCount: number }) => (
  <Tabs value={value - 1} variant={variant}>
    {Array(tabCount)
      .fill(null)
      .map((_, index: number) => (
        // eslint-disable-next-line react/no-array-index-key
        <Tab key={index} label={`Tab ${index + 1}`} value={index}>
          <div>Tab{index + 1} Content</div>
        </Tab>
      ))}
  </Tabs>
);

export const Primary = Template.bind({});
Primary.args = {
  tabCount: 2,
  value: 1,
};

export const FullWidth = Template.bind({});
FullWidth.args = {
  tabCount: 2,
  value: 1,
  variant: "fullWidth",
};
