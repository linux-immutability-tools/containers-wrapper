package types

/*	License: GPLv3
	Authors:
		Mirko Brombin <mirko@fabricators.ltd>
		Lit Contributors <https://github.com/linux-immutability-tools/>
	Copyright: 2023
	Description:
		Containers Wrapper is a Go library that provides a convenient
		and unified interface for interacting with container engines
		such as Docker, Podman, and Containerd.
*/

import (
	"errors"
)

var (
	// Container Engine errors

	ErrNoContainerEngine       = errors.New("[ce] [FAIL]: no container engine defined")
	ErrContainerEngineNotFound = errors.New("[ce] [FAIL]: defined container engine not found. \n\t\t → Is it installed?")

	// Images errors

	ErrImagesGenericFailure         = errors.New("[images] [FAIL]: for unknown reason")
	ErrImagesEmptyID                = errors.New("[images] [FAIL]: empty image id")
	ErrImagesBadFilter              = errors.New("[images] [FAIL]: invalid filter specified")
	ErrImagesEmptyTarball           = errors.New("[images] [FAIL]: empty tarball")
	ErrImagesTarballNotFound        = errors.New("[images] [FAIL]: tarball not found or not accessible")
	ErrImagesEmptyDestination       = errors.New("[images] [FAIL]: empty destination")
	ErrImagesNotFound               = errors.New("[images] [FAIL]: image not found")
	ErrImagesDestinationNotFound    = errors.New("[images] [FAIL]: destination not found or not accessible")
	ErrImagesHasMalformedName       = errors.New("[images] [FAIL]: image has malformed name. \n\t\t → The format should be <repository>:<tag>")
	ErrImagesMultipleResults        = errors.New("[images] [FAIL]: multiple results found. \n\t\t → This should never happen, what's wrong with your setup?")
	ErrImagesContainerfileNotFound  = errors.New("[images] [FAIL]: Containerfile not found or not accessible")
	ErrImagesContainerfileBadFormat = errors.New("[images] [FAIL]: Containerfile has bad format")

	// Containers errors

	ErrContainersGenericFailure  = errors.New("containers [FAIL]: for unknown reason")
	ErrContainersBadFilter       = errors.New("containers [FAIL]: invalid filter specified")
	ErrContainersNotFound        = errors.New("containers [FAIL]: container not found")
	ErrContainerNameAlreadyInUse = errors.New("containers [FAIL]: container name already in use")
)
