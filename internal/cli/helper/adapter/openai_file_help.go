package adapter

import "fmt"

func (h *OpenAiFileAdapter) FileListHelpInfo(adp *OpenAiAdapter) string {
	return "You can repeat flags to provide more than one filter, e.g.:\n" +
		fmt.Sprintf(
			"  %s --%s %s --%s %s --%s %s",
			h.command.Name(),
			PurposeFlagKey, adp.enumset.Purpose().Codes.FineTune,
			PurposeFlagKey, PurposeExample,
			CreatedAfterFlagKey, DateExample,
		)
}

func (h *OpenAiFileAdapter) FileDeleteHelpInfo(adp *OpenAiAdapter) string {
	return h.FileListHelpInfo(adp)
}

func (*OpenAiFileAdapter) FileUploadHelpInfo(adp *OpenAiAdapter) string {
	return "Each argument can optionally have its own file purpose" +
		" by appending the suffix ':purpose'\n" +
		fmt.Sprintf(
			"Defaults to '%s' if --%s is not provided.",
			adp.enumset.Purpose().Default().Code,
			DefaultPurposeFlagKey,
		)
}
