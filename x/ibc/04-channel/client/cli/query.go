package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/ibc/04-channel/client/utils"
	"github.com/cosmos/cosmos-sdk/x/ibc/04-channel/types"
	host "github.com/cosmos/cosmos-sdk/x/ibc/24-host"
)

const (
	flagSequences = "sequences"
)

// GetCmdQueryChannels defines the command to query all the channels ends
// that this chain mantains.
func GetCmdQueryChannels() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "channels",
		Short:   "Query all channels",
		Long:    "Query all channels from a chain",
		Example: fmt.Sprintf("%s query %s %s channels", version.AppName, host.ModuleName, types.SubModuleName),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			req := &types.QueryChannelsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.Channels(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "channels")

	return cmd
}

// GetCmdQueryChannel defines the command to query a channel end
func GetCmdQueryChannel() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "end [port-id] [channel-id]",
		Short: "Query a channel end",
		Long:  "Query an IBC channel end from a port and channel identifiers",
		Example: fmt.Sprintf(
			"%s query %s %s end [port-id] [channel-id]", version.AppName, host.ModuleName, types.SubModuleName,
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			portID := args[0]
			channelID := args[1]
			prove, _ := cmd.Flags().GetBool(flags.FlagProve)

			channelRes, err := utils.QueryChannel(clientCtx, portID, channelID, prove)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(channelRes)
		},
	}

	cmd.Flags().Bool(flags.FlagProve, true, "show proofs for the query results")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryConnectionChannels defines the command to query all the channels associated with a
// connection
func GetCmdQueryConnectionChannels() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "connections [connection-id]",
		Short:   "Query all channels associated with a connection",
		Long:    "Query all channels associated with a connection",
		Example: fmt.Sprintf("%s query %s %s connections [connection-id]", version.AppName, host.ModuleName, types.SubModuleName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			req := &types.QueryConnectionChannelsRequest{
				Connection: args[0],
				Pagination: pageReq,
			}

			res, err := queryClient.ConnectionChannels(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "channels associated with a connection")

	return cmd
}

// GetCmdQueryChannelClientState defines the command to query a client state from a channel
func GetCmdQueryChannelClientState() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "client-state [port-id] [channel-id]",
		Short:   "Query the client state associated with a channel",
		Long:    "Query the client state associated with a channel, by providing its port and channel identifiers.",
		Example: fmt.Sprintf("%s query ibc channel client-state [port-id] [channel-id]", version.AppName),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			portID := args[0]
			channelID := args[1]

			res, err := utils.QueryChannelClientState(clientCtx, portID, channelID, false)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res.IdentifiedClientState)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryPacketCommitments defines the command to query all packet commitments associated with
// a channel
func GetCmdQueryPacketCommitments() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "packet-commitments [port-id] [channel-id]",
		Short:   "Query all packet commitments associated with a channel",
		Long:    "Query all packet commitments associated with a channel",
		Example: fmt.Sprintf("%s query %s %s packet-commitments [port-id] [channel-id]", version.AppName, host.ModuleName, types.SubModuleName),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			req := &types.QueryPacketCommitmentsRequest{
				PortId:     args[0],
				ChannelId:  args[1],
				Pagination: pageReq,
			}

			res, err := queryClient.PacketCommitments(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "packet commitments associated with a channel")

	return cmd
}

// GetCmdQueryPacketCommitment defines the command to query a channel end
func GetCmdQueryPacketCommitment() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "packet-commitment [port-id] [channel-id] [sequence]",
		Short: "Query a packet commitment",
		Long:  "Query a packet commitment",
		Example: fmt.Sprintf(
			"%s query %s %s end [port-id] [channel-id]", version.AppName, host.ModuleName, types.SubModuleName,
		),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			portID := args[0]
			channelID := args[1]
			prove, _ := cmd.Flags().GetBool(flags.FlagProve)

			seq, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}

			res, err := utils.QueryPacketCommitment(clientCtx, portID, channelID, seq, prove)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}

	cmd.Flags().Bool(flags.FlagProve, true, "show proofs for the query results")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryUnreceivedPackets defines the command to query all the unreceived
// packets on the receiving chain
func GetCmdQueryUnreceivedPackets() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unreceived-packets [port-id] [channel-id]",
		Short: "Query all the unreceived packets associated with a channel",
		Long: `Determine if a packet, given a list of packet commitment sequences, is unreceived.

The return value represents:
- Unreceived packet commitments: no acknowledgement exists on receiving chain for the given packet commitment sequence on sending chain.
`,
		Example: fmt.Sprintf("%s query %s %s unreceived-packets [port-id] [channel-id] --sequences=1,2,3", version.AppName, host.ModuleName, types.SubModuleName),
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			seqSlice, err := cmd.Flags().GetInt64Slice(flagSequences)
			if err != nil {
				return err
			}

			seqs := make([]uint64, len(seqSlice))
			for i := range seqSlice {
				seqs[i] = uint64(seqSlice[i])
			}

			req := &types.QueryUnreceivedPacketsRequest{
				PortId:                    args[0],
				ChannelId:                 args[1],
				PacketCommitmentSequences: seqs,
			}

			res, err := queryClient.UnreceivedPackets(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}

	cmd.Flags().Int64Slice(flagSequences, []int64{}, "comma separated list of packet sequence numbers")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryUnrelayedAcks defines the command to query all the unrelayed acks on the original sending chain
func GetCmdQueryUnrelayedAcks() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unrelayed-acks [port-id] [channel-id]",
		Short: "Query all the unrelayed acks associated with a channel",
		Long: `Given a list of packet commitment sequences from counterparty, determine if an ack on executing chain has not been relayed to counterparty.

The return value represents:
- Unrelayed packet acknowledgement: packet commitment exists on original sending chain and ack exists on receiving (executing) chain.
`,
		Example: fmt.Sprintf("%s query %s %s unrelayed-acks [port-id] [channel-id] --sequences=1,2,3", version.AppName, host.ModuleName, types.SubModuleName),
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			seqSlice, err := cmd.Flags().GetInt64Slice(flagSequences)
			if err != nil {
				return err
			}

			seqs := make([]uint64, len(seqSlice))
			for i := range seqSlice {
				seqs[i] = uint64(seqSlice[i])
			}

			req := &types.QueryUnrelayedAcksRequest{
				PortId:                    args[0],
				ChannelId:                 args[1],
				PacketCommitmentSequences: seqs,
			}

			res, err := queryClient.UnrelayedAcks(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}

	cmd.Flags().Int64Slice(flagSequences, []int64{}, "comma separated list of packet sequence numbers")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryNextSequenceReceive defines the command to query a next receive sequence for a given channel
func GetCmdQueryNextSequenceReceive() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "next-sequence-receive [port-id] [channel-id]",
		Short: "Query a next receive sequence",
		Long:  "Query the next receive sequence for a given channel",
		Example: fmt.Sprintf(
			"%s query %s %s next-sequence-receive [port-id] [channel-id]", version.AppName, host.ModuleName, types.SubModuleName,
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			portID := args[0]
			channelID := args[1]
			prove, _ := cmd.Flags().GetBool(flags.FlagProve)

			sequenceRes, err := utils.QueryNextSequenceReceive(clientCtx, portID, channelID, prove)
			if err != nil {
				return err
			}

			clientCtx = clientCtx.WithHeight(int64(sequenceRes.ProofHeight.EpochHeight))
			return clientCtx.PrintOutput(sequenceRes)
		},
	}

	cmd.Flags().Bool(flags.FlagProve, true, "show proofs for the query results")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
