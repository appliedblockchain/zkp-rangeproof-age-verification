import com.ing.blockchain.zk.components.SecretOrderGroupGenerator;
import com.ing.blockchain.zk.dto.*;
import com.ing.blockchain.zk.util.ExportUtil;

import javax.xml.bind.DatatypeConverter;

import java.math.BigInteger;

// sample code for zero knowledge age verification - government/kyc-provider provides a proof (TTPMessage - trusted third party message)

public class ZKRPSample {
    static String age = "19";

    /**
     * Create a secret order group (for use in generating a commitment). A larger bit length is more secure but less efficient
     * @return SecretOrderGroup
     */
    public static SecretOrderGroup generateSecretGroup() {
        return new SecretOrderGroupGenerator(512).generate();
    }

    /**
     * Govt or trusted authority generates a commitment on a value (in this case a user's age)
     * @return TTPMessage
     */
    public static TTPMessage generateTTPMessage () {
        BigInteger commitmentValue = new BigInteger(age);

        SecretOrderGroup secretGroup = generateSecretGroup();

        // TODO: this message needs to be signed by the govt so that the verifier can check the authenticity of the
        // proof via EC signature verification (JS impl in https://github.com/appliedblockchain/zkp-integer-threshold)
        TTPMessage ttpMessage = Govt.createTrustedMessage(commitmentValue, secretGroup);

        return ttpMessage;
    }

    public static void main(String[] args) {
        TTPMessage ttpMessage = generateTTPMessage();

        Commitment commitment = ttpMessage.getCommitment();

        ClosedRange ageRange = Config.getAgeRange();

        BoudotRangeProof ageProof = User.generateRangeProof(ttpMessage, ageRange);

        // Print commitment
        Utils.print("0x" + DatatypeConverter.printHexBinary(ExportUtil.exportForEVM(commitment)));

        // Print delimiter (used to split the commitment and proof when calling the .jar via terminal)
        Utils.print("|");

        // Print proof
        Utils.print("0x" + DatatypeConverter.printHexBinary(ExportUtil.exportForEVM(ageProof, commitment, ageRange)));
    }
}
