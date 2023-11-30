import type { ClutchColors, ComponentState, ThemeVariant } from "./types";

export const LIGHT_COLORS: ClutchColors = {
  neutral: {
    50: "#F8F8F9",
    100: "#F0F1F3",
    200: "#DBDBE0",
    300: "#C2C3CB",
    400: "#9E9FAC",
    500: "#868798",
    600: "#56586E",
    700: "#3D4059",
    800: "#252845",
    900: "#0D1030",
    A100: "#FFFFFF",
    A200: "#FFFFFF",
    A400: "#FFFFFF",
    A700: "#FFFFFF",
  },
  blue: {
    50: "#F9F9FE",
    100: "#F5F6FD",
    200: "#EBEDFB",
    300: "#D7DAF6",
    400: "#C2C8F2",
    500: "#727FE1",
    600: "#3548D4",
    700: "#1629B9",
    800: "#0A1CA6",
    900: "#011082",
    A100: "#FFFFFF",
    A200: "#FFFFFF",
    A400: "#FFFFFF",
    A700: "#FFFFFF",
  },
  green: {
    50: "#E5FCE8",
    100: "#CBF7CF",
    200: "#ACF2B2",
    300: "#6CD47A",
    400: "#4AB958",
    500: "#32A140",
    600: "#1C872A",
    700: "#106E1D",
    800: "#086515",
    900: "#02590E",
    A100: "#FFFFFF",
    A200: "#FFFFFF",
    A400: "#FFFFFF",
    A700: "#FFFFFF",
  },
  amber: {
    50: "#FFFBEB",
    100: "#FEF3C7",
    200: "#FDE68A",
    300: "#FCD34D",
    400: "#FBBF24",
    500: "#F59E0B",
    600: "#D97706",
    700: "#B45309",
    800: "#92400E",
    900: "#78350F",
    A100: "#FFFFFF",
    A200: "#FFFFFF",
    A400: "#FFFFFF",
    A700: "#FFFFFF",
  },
  red: {
    50: "#FFF4F3",
    100: "#FDDCD2",
    200: "#F6A996",
    300: "#F3886E",
    400: "#F26E50",
    500: "#F15534",
    600: "#CA4428",
    700: "#A1301C",
    800: "#792111",
    900: "#571608",
    A100: "#FFFFFF",
    A200: "#FFFFFF",
    A400: "#FFFFFF",
    A700: "#FFFFFF",
  },
};

export const DARK_COLORS: ClutchColors = {
  neutral: {
    50: "#232542",
    100: "#30324E",
    200: "#3E4059",
    300: "#55566D",
    400: "#77788A",
    500: "#8D8E9E",
    600: "#A4A5B1",
    700: "#D2D2D8",
    800: "#E8E8EB",
    900: "#FFFFFF",
    A100: "#FFFFFF",
    A200: "#FFFFFF",
    A400: "#FFFFFF",
    A700: "#FFFFFF",
  },
  blue: {
    50: "#090A42",
    100: "#050656",
    200: "#0A086B",
    300: "#13199D",
    400: "#2A4FF6",
    500: "#4281F6",
    600: "#5AABF6",
    700: "#8CC4F8",
    800: "#C2E1FE",
    900: "#DCECFB",
    A100: "#FFFFFF",
    A200: "#FFFFFF",
    A400: "#FFFFFF",
    A700: "#FFFFFF",
  },
  green: {
    50: "#002C05",
    100: "#084713",
    200: "#145F1D",
    300: "#217A2A",
    400: "#2D9638",
    500: "#54B45B",
    600: "#73C178",
    700: "#9CD29E",
    800: "#C3E4C4",
    900: "#E6F4E7",
    A100: "#FFFFFF",
    A200: "#FFFFFF",
    A400: "#FFFFFF",
    A700: "#FFFFFF",
  },
  amber: {
    50: "#352215",
    100: "#7E3E02",
    200: "#985304",
    300: "#B26A08",
    400: "#CC820D",
    500: "#E69F2A",
    600: "#EAB04E",
    700: "#EFC67F",
    800: "#F6DCB1",
    900: "#FBF1E0",
    A100: "#FFFFFF",
    A200: "#FFFFFF",
    A400: "#FFFFFF",
    A700: "#FFFFFF",
  },
  red: {
    50: "#501306",
    100: "#621809",
    200: "#751E0D",
    300: "#972814",
    400: "#B3351E",
    500: "#C8482C",
    600: "#EB7A60",
    700: "#F4A08A",
    800: "#F0B9AB",
    900: "#FBE8E7",
    A100: "#FFFFFF",
    A200: "#FFFFFF",
    A400: "#FFFFFF",
    A700: "#FFFFFF",
  },
};

export const STATE_OPACITY: { [key in ComponentState]: number } = {
  hover: 0.5,
  focused: 0.1,
  pressed: 0.15,
  selected: 0.12,
  disabled: 0.5,
};

const clutchColors = (variant: ThemeVariant) => {
  const colors = variant === "light" ? LIGHT_COLORS : DARK_COLORS;
  return {
    ...colors,
  };
};

export { clutchColors };
