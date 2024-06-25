package nex

import (
	"github.com/PretendoNetwork/dr-luigi/globals"
	"github.com/PretendoNetwork/dr-luigi/database"
	matchmake_extension "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension"
	common_matchmake_extension "github.com/PretendoNetwork/nex-protocols-common-go/v2/matchmake-extension"
	match_making "github.com/PretendoNetwork/nex-protocols-go/v2/match-making"
	common_match_making "github.com/PretendoNetwork/nex-protocols-common-go/v2/match-making"
	match_making_ext "github.com/PretendoNetwork/nex-protocols-go/v2/match-making-ext"
	common_match_making_ext "github.com/PretendoNetwork/nex-protocols-common-go/v2/match-making-ext"
	ranking "github.com/PretendoNetwork/nex-protocols-go/v2/ranking"
	common_ranking "github.com/PretendoNetwork/nex-protocols-common-go/v2/ranking"
	nat_traversal "github.com/PretendoNetwork/nex-protocols-go/v2/nat-traversal"
	common_nat_traversal "github.com/PretendoNetwork/nex-protocols-common-go/v2/nat-traversal"
	secure "github.com/PretendoNetwork/nex-protocols-go/v2/secure-connection"
	common_secure "github.com/PretendoNetwork/nex-protocols-common-go/v2/secure-connection"
	utility "github.com/PretendoNetwork/nex-protocols-go/v2/utility"
	common_utility "github.com/PretendoNetwork/nex-protocols-common-go/v2/utility"

	nex_matchmake_extension_common "github.com/PretendoNetwork/dr-luigi/nex/matchmake-extension/common"
)

func registerCommonSecureServerProtocols() {
	secureProtocol := secure.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(secureProtocol)
	common_secure.NewCommonProtocol(secureProtocol)

	natTraversalProtocol := nat_traversal.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(natTraversalProtocol)
	common_nat_traversal.NewCommonProtocol(natTraversalProtocol)

	matchMakingProtocol := match_making.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(matchMakingProtocol)
	common_match_making.NewCommonProtocol(matchMakingProtocol)

	matchMakingExtProtocol := match_making_ext.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(matchMakingExtProtocol)
	common_match_making_ext.NewCommonProtocol(matchMakingExtProtocol)

	matchmakeExtensionProtocol := matchmake_extension.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(matchmakeExtensionProtocol)
	commonMatchmakeExtensionProtocol := common_matchmake_extension.NewCommonProtocol(matchmakeExtensionProtocol)

	commonMatchmakeExtensionProtocol.CleanupSearchMatchmakeSession = nex_matchmake_extension_common.CleanupSearchMatchmakeSession
	commonMatchmakeExtensionProtocol.GetUserFriendPIDs(nex_matchmake_extension_common.GetUserFriendPIDs)

	rankingProtocol := ranking.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(rankingProtocol)
	ranking_protocol := common_ranking.NewCommonProtocol(rankingProtocol)
	ranking_protocol.GetRankingsAndCountByCategoryAndRankingOrderParam = database.GetRankingsAndCountByCategoryAndRankingOrderParam
	ranking_protocol.InsertRankingByPIDAndRankingScoreData = database.InsertRankingByPIDAndRankingScoreData
	ranking_protocol.UploadCommonData = database.UploadCommonData

	utilityProtocol := utility.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(utilityProtocol)
	common_utility.NewCommonProtocol(utilityProtocol)
}
