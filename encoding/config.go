// Copyright 2022 Match Foundation
// This file is part of the Match Network packages.
//
// Match is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Match packages are distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Match packages. If not, see https://github.com/matchain/match/blob/main/LICENSE
package encoding

import (
	amino "github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"

	enccodec "github.com/matchain/match/encoding/codec"
)

// MakeConfig creates an EncodingConfig for testing
func MakeConfig(mb module.BasicManager) params.EncodingConfig {
	cdc := amino.NewLegacyAmino()
	interfaceRegistry := types.NewInterfaceRegistry()
	codec := amino.NewProtoCodec(interfaceRegistry)

	encodingConfig := params.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Codec:             codec,
		TxConfig:          tx.NewTxConfig(codec, tx.DefaultSignModes),
		Amino:             cdc,
	}

	enccodec.RegisterLegacyAminoCodec(encodingConfig.Amino)
	mb.RegisterLegacyAminoCodec(encodingConfig.Amino)
	enccodec.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	mb.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}
