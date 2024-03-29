import * as React from "react";
import styled from "@emotion/styled";
import type { Theme } from "@mui/material";
import { alpha } from "@mui/system";

const HorizontalRuleBase = ({ children, ...props }: HorizontalRuleProps) => (
  <div {...props}>
    <div className="line">
      <span />
    </div>
    {React.Children.count(children) > 0 && <div className="content">{children}</div>}
    <div className="line">
      <span />
    </div>
  </div>
);

export type HorizontalRuleProps = {
  children: React.ReactNode;
};

const StyledHorizontalRule = styled(HorizontalRuleBase)(({ theme }: { theme: Theme }) => ({
  alignItems: "center",
  display: "flex",
  flexDirection: "row",
  width: "100%",
  margin: "16px 0px",

  ".line": {
    flex: "1 1 auto",
  },

  ".line > span": {
    display: "block",
    borderTop: `1px solid ${alpha(theme.palette.secondary[900], 0.12)}`,
  },

  ".content": {
    padding: "0 16px",
    fontWeight: "bold",
    fontSize: "14px",
    color: alpha(theme.palette.secondary[900], 0.38),
    textTransform: "uppercase",
    display: "inline-flex",
    alignItems: "center",
  },
}));

export const HorizontalRule = ({ children }: HorizontalRuleProps) => (
  <StyledHorizontalRule>{children}</StyledHorizontalRule>
);

export default HorizontalRule;
