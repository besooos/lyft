import * as React from "react";
import styled from "@emotion/styled";
import CloseIcon from "@mui/icons-material/Close";
import type { DialogProps as MuiDialogProps } from "@mui/material";
import {
  alpha,
  Dialog as MuiDialog,
  DialogActions as MuiDialogActions,
  DialogContent as MuiDialogContent,
  DialogTitle as MuiDialogTitle,
  IconButton as MuiIconButton,
  Paper,
} from "@mui/material";

const DialogPaper = styled(Paper)(({ theme }) => ({
  border: `1px solid ${alpha(theme.palette.secondary[900], 0.1)}`,
  boxShadow: "0px 10px 24px rgba(35, 48, 143, 0.3)",
  boxSizing: "border-box",
  backgroundColor: theme.palette.contrastColor,
  width: "max-content",
  maxWidth: "75vw",
}));

const DialogTitle = styled(MuiDialogTitle)(({ theme }) => ({
  display: "flex",
  justifyContent: "space-between",
  fontSize: "20px",
  padding: "12px 12px 0 32px",
  fontWeight: 500,
  color: theme.palette.secondary[900],
}));

const DialogTitleText = styled.div({
  padding: "14px 0 0 0",
});

const IconButton = styled(MuiIconButton)(({ theme }) => ({
  height: "12px",
  width: "12px",
  color: theme.palette.secondary[900],
}));

const DialogContent = styled(MuiDialogContent)(({ theme }) => ({
  padding: "16px 32px 32px 32px",
  fontSize: "16px",
  fontWeight: 400,
  color: alpha(theme.palette.secondary[900], 0.6),
  "> *": {
    margin: "16px 0 0 0",
  },
  overflowWrap: "break-word",
}));

const DialogActions = styled(MuiDialogActions)(({ theme }) => ({
  borderTop: `1px solid ${alpha(theme.palette.secondary[900], 0.12)}`,
  padding: "0 8px",
  "> *": {
    margin: "16px 8px 16px 8px",
  },
}));

export interface DialogContentProps {}
export interface DialogActionsProps {}
export type DialogCloseReasons = "closeButtonClick" | "escapeKeyDown" | "backdropClick";
export interface DialogProps extends Pick<MuiDialogProps, "open"> {
  title: string;
  children:
    | React.ReactElement<DialogContentProps>
    | React.ReactElement<DialogContentProps>[]
    | React.ReactElement<DialogActionsProps>
    | React.ReactElement<DialogActionsProps>[];
  onClose: (event?: object, reason?: DialogCloseReasons) => void;
}

const Dialog = ({ title, children, open, onClose }: DialogProps) => (
  <MuiDialog PaperComponent={DialogPaper} open={open} onClose={onClose} maxWidth={false}>
    <DialogTitle>
      <DialogTitleText>{title}</DialogTitleText>
      <IconButton onClick={e => onClose(e, "closeButtonClick")} size="large">
        <CloseIcon />
      </IconButton>
    </DialogTitle>
    {children}
  </MuiDialog>
);

export { Dialog, DialogActions, DialogContent };
