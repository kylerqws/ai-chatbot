package adapter

import "fmt"

func (h *OpenAiJobAdapter) JobListHelpInfo(adp *OpenAiAdapter) string {
	return "You can repeat flags to provide more than one filter, e.g.:\n" +
		fmt.Sprintf(
			"  %s --%s %s --%s %s --%s %s",
			h.command.Name(),
			ModelFlagKey, adp.enumset.Model().Codes.GPT35,
			ModelFlagKey, ModelExample,
			CreatedAfterFlagKey, DateExample,
		)
}

func (h *OpenAiJobAdapter) JobCancelHelpInfo(adp *OpenAiAdapter) string {
	return h.JobListHelpInfo(adp)
}

func (*OpenAiJobAdapter) JobCreateHelpInfo(adp *OpenAiAdapter) string {
	return "Each argument can optionally have its own validation file and model" +
		" by appending the suffixes ':validation-file-id' and ':model'\n" +
		fmt.Sprintf(
			"Defaults to '%s' if --%s is not provided.",
			adp.enumset.Model().Default().Code,
			DefaultModelFlagKey,
		)
}
