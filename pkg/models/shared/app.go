// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type App struct {
	// The command line arguments of the executable.
	Args []string `json:"args,omitempty"`
	// The working directory of the executable.
	Cwd *string `json:"cwd,omitempty"`
	// The name of the executable.
	Executable string `json:"executable"`
	// The directory of the executable.
	ExecutablePath *string `json:"executablePath,omitempty"`
	ID             *string `json:"id,omitempty"`
	// pods that match one of the given labels
	Labels []Label `json:"labels,omitempty"`
	// The unique name of the App (as it will appear in the Secure Application UI).
	Name string `json:"name"`
	// The process name.
	ProcessName *string `json:"processName,omitempty"`
	// Type of the App.
	Type string `json:"type"`
}

func (o *App) GetArgs() []string {
	if o == nil {
		return nil
	}
	return o.Args
}

func (o *App) GetCwd() *string {
	if o == nil {
		return nil
	}
	return o.Cwd
}

func (o *App) GetExecutable() string {
	if o == nil {
		return ""
	}
	return o.Executable
}

func (o *App) GetExecutablePath() *string {
	if o == nil {
		return nil
	}
	return o.ExecutablePath
}

func (o *App) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *App) GetLabels() []Label {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *App) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *App) GetProcessName() *string {
	if o == nil {
		return nil
	}
	return o.ProcessName
}

func (o *App) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}
