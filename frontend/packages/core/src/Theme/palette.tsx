import type { PaletteOptions as MuiPaletteOptions } from "@mui/material/styles";
import { alpha, TypeText } from "@mui/material/styles";

import { DARK_COLORS, LIGHT_COLORS } from "./colors";
import type { ClutchColors, ThemeVariant } from "./types";

interface PaletteOptions extends MuiPaletteOptions {
  type: ThemeVariant;
}

const lightText: Partial<TypeText> = {
  primary: LIGHT_COLORS.neutral[900],
  secondary: alpha(LIGHT_COLORS.neutral[900], 0.65),
  // tertiary
  // inverse
};

const darkText: Partial<TypeText> = {
  primary: alpha(DARK_COLORS.neutral[900], 0.9),
  secondary: alpha(DARK_COLORS.neutral[900], 0.75),
  // tertiary
  // inverse
};

const palette = (variant: ThemeVariant): PaletteOptions => {
  const isLightMode = variant === "light";
  const color = (isLightMode ? LIGHT_COLORS : DARK_COLORS) as ClutchColors;

  // TODO: add all clutch colors to "common colors"
  return {
    type: variant,
    mode: variant,
    primary: color.blue,
    secondary: color.neutral,
    error: color.red,
    warning: color.amber,
    info: color.blue,
    success: color.green,
    grey: color.neutral,
    background: {
      default: color.blue[50],
      paper: isLightMode ? color.neutral.A100 : "#1c1e3c",
    },
    text: isLightMode ? lightText : darkText,
  };
};

export default palette;
