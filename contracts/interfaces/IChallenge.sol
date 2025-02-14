// SPDX-License-Identifier: GPL v3
pragma solidity ^0.8.0;

import "../libraries/Types.sol";

interface IChallenge {
    //the info of rv32 system info.
    struct SystemInfo {
        Types.StateInfo stateInfo;
        //systemEndState index,must > 1.
        uint128 endStep;
        //systemStartState calculated by Executor.
        bytes32 systemStartState;
        bytes32 systemEndState;
    }

    enum ChallengeStage {
        Uninitialized,
        //challenge game started.
        Started,
        // challenge game initialized,and now need to find out one step.
        Running,
        // one step find out and challenger win, or proposer win, game over.
        Finished
    }
    //challenger whether need to claim dispute proposer's staking.
    enum ClaimStatus {
        //when game finished and challenger win, they have to claim the payback.
        UnClaimed,
        //claim over.
        Over
    }

    function minChallengerDeposit() external view returns (uint256);

    /**
     * @dev Create challenge by challengeFactory.guarantee the info provided true.
     * @param _systemStartState System initial state of program, calculated by executor.
     * @param _creator Challenger who start challenge.
     * @param _proposerTimeLimit After how much l1 block, the proposer expired.
     * @param _stateInfo StateInfo contains the challenged block info, already confirmed by challengeFactory
     * @param _minChallengerDeposit floor price for challenge to engage challenge game
     * @notice revert when transfer failed
     */
    function create(
        bytes32 _systemStartState,
        address _creator,
        uint256 _proposerTimeLimit,
        Types.StateInfo memory _stateInfo,
        uint256 _minChallengerDeposit
    ) external;

    event ChallengeInitialized(uint128 _systemEndStep, bytes32 _midSystemState);

    /**
     * @dev Initialize challenge info, provide endStep and mid system state of program.
     * @param _endStep End step index of system state of program,must larger than 1.
     * @param _systemEndState End system state, 0 is illegal, and end state must be "correct"(the program is halt, and the output is consistent with outputRoot).
     * @param _midSystemState Mid state root of system,0 is illegal.
     */
    function initialize(
        uint64 _endStep,
        bytes32 _systemEndState,
        bytes32 _midSystemState
    ) external;

    event MidStateRevealed(uint256[] nodeKeys, bytes32[] stateRoots);

    /**
     * @dev Proposer reveal the node's midRoots, he can reveal in advance.
     * @param _nodeKeys The revealed keys in disputeTree in order.
     * @param _stateRoots The revealed mid state roots of above nodeKey,0 state root is illegal.
     * @notice Revert if provide empty slice, slice's length different or not revealed in time；or stateHash is equal to 0;or
     * attempt to re-reveal exist midState.
     */
    function revealMidStates(uint256[] calldata _nodeKeys, bytes32[] calldata _stateRoots) external;

    event ProposerTimeout(uint256 nodeKey);

    /**
     * @dev proposer only need to do 2 kinds of action:
     * 1: initialize challenge info after challenge started.
     * 2: provide mid state when challenge running.
     * if one of above timeout, anyone can call this to end challenge game.
     * @param _nodeKey If proposer timeout doing 1st situation, node key can easily set to 0.
     * Otherwise, in 2nd situation, _nodeKey is the disputeNode key which proposer didn't reveal in expireTime.
     */
    function proposerTimeout(uint256 _nodeKey) external;

    event DisputeBranchSelected(address indexed challenger, uint256[] nodeKey, uint256 expireAfterBlock);

    /**
     * @dev Anyone has deposited in this challengeGame can select one branch in dispute tree.which means selected dispute
     * nodes' start system state is right,and the nodes' end system state is wrong.
     * @param _parentNodeKey The parent node key in disputeTree.When we select a dispute branch, it must derived from exist larger disputeNode, we call it parent node
     * i.e. A node present 0->4 stateTransition, B node present 0->2 stateTransition,0->2 is driven by 0->4,so A is parent node.
     * @param _isLeft Select whether left or right bisection child of parent node.i.e. parent node present 0->4 transition, left child is
     * 0->2, right child is 2->4.
     * @notice Revert if chose more than one branch; or parent node not exist;or has no provided mid state;
     * or one step node is the parent.
     */
    function selectDisputeBranch(uint256[] calldata _parentNodeKey, bool[] calldata _isLeft) external;

    event OneStepTransition(uint256 startStep, bytes32 revealedRoot, bytes32 executedRoot);

    /**
     * @dev Anyone can verify oneStepRuns in challenge period.if the result system state of one step is not same as proposer claimed. challenger win roll back the block and game over.It ensures:
     * 1: in interactive period.
     * 2: proposer is wrong.
     * 3: roll back the wrong block
     * @param _leafNodeKey is one step node key in dispute tree.
     * @notice Revert if node is not oneStepNode.
     */
    function execOneStepTransition(uint256 _leafNodeKey) external;

    event ProposerWin(address _winner, uint256 _amount);

    /**
     * @dev The only way can proposer win.This allows unfinished challenge game over after passing the challenge time.
     */
    function claimProposerWin() external;

    /**
     * @dev Challenger can get reward by this way if there exist only one branch.Otherwise, now just transfer to DAO.
     * @param _challenger Which challenger claims the payback
     * @param _stateInfo StateInfo to provide the New stateInfo
     */
    function claimChallengerWin(address _challenger, Types.StateInfo memory _stateInfo) external;
}
